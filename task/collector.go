package task

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/its-my-data/doubak/proto"
	"github.com/its-my-data/doubak/util"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// TODO: use a separate library for URLs.

const DoubanURL = "https://www.douban.com/"
const PeopleURL = DoubanURL + "people/"

// Collector contains the information used by the collector.
type Collector struct {
	user       string
	categories []string
	outputDir  string
}

// NewCollector returns a new collector task and initialise it.
func NewCollector(userName *string, categories []string) *Collector {
	return &Collector{
		user:       *userName,
		categories: categories,
	}
}

// Precheck validates the flags.
func (task *Collector) Precheck() error {
	// Initialize the top most directory for Collector.
	if path, err := util.GetPathWithCreation(util.CollectorPathPrefix); err != nil {
		return err
	} else {
		task.outputDir = path
	}
	log.Println("New output path saved:", task.outputDir)

	// Check user existence.
	exists := true
	cu := colly.NewCollector()
	// TODO: remove.
	//cu.OnResponse(func(r *colly.Response) {
	//	log.Println(string(r.Body))
	//})
	//cu.OnHTML("li", func(e *colly.HTMLElement) {
	//	log.Println(e.Text)
	//})
	cu.SetRequestTimeout(5 * time.Minute)
	cu.OnError(func(r *colly.Response, err error) {
		exists = false
		t := string(r.Body)
		if strings.Contains(t, "页面不存在") {
			// Real user not found.
		} else {
			// Usually caused by timeouts, but just in case.
			log.Println("Request URL:", r.Request.URL, "\nError:", err)
			log.Println("Unknown error with request body: ", string(r.Body))
		}
	})

	// Error handled separately.
	_ = cu.Visit(PeopleURL + task.user + "/")

	if !exists {
		return errors.New("user '" + task.user + "' does not exist")
	}
	return nil
}

// Execute starts the collection.
func (task *Collector) Execute() error {
	for _, c := range task.categories {
		switch c {
		case proto.Category_broadcast.String():
			task.crawlBroadcasts()
		case proto.Category_book.String():
			task.crawlBooks()
		case proto.Category_movie.String():
			task.crawlMovies()
		case proto.Category_game.String():
			task.crawlGames()
		default:
			return errors.New("Category not implemented " + c)
		}
	}
	return nil
}

func (task *Collector) crawlBroadcasts() error {
	// TODO: need to be global or class-private.
	// These can be hacked to resume a progress.
	timePrefix := time.Now().Local().Format("20060102.1504")
	page := 1

	q := util.NewQueue()
	c := util.NewColly()
	c.OnResponse(func(r *colly.Response) {
		if err := task.saveResponse(r, proto.Category_broadcast, timePrefix, page); err != nil {
			log.Println(err.Error())
		}

		body := string(r.Body)
		util.FailIfNeedLogin(&body)

		// Prepare for the next request.
		// Note that the number of broadcasts in each page somehow don't equal.
		// Therefore, I have to get at least an empty status page file.
		broadcastCount := strings.Count(body, "\"status-item\"")
		log.Println("Found", broadcastCount, "broadcasts/status.")
		if broadcastCount != 0 {
			page++
			url := PeopleURL + task.user + "/statuses?p=" + strconv.Itoa(page)
			q.AddURL(url)
			log.Printf("Added URL: %s. (Followed by sleeping.)\n", url)
			time.Sleep(3 * time.Second)
		} else {
			log.Printf("All done with broadcast count %d (in page %d).\n", broadcastCount, page)
		}
	})

	// TODO: need a retry queue (either Requests, or go routines).
	q.AddURL(PeopleURL + task.user + "/statuses?p=" + strconv.Itoa(page))

	return q.Run(c)
}

func (task *Collector) crawlBooks() error {
	// TODO: update the implementation.
	return errors.New("update the implementation")
}

func (task *Collector) crawlMovies() error {
	// TODO: update the implementation.
	return errors.New("update the implementation")
}

func (task *Collector) crawlGames() error {
	// TODO: update the implementation.
	return errors.New("update the implementation")
}

// TODO: implement more crawlers.

func (task *Collector) saveResponse(r *colly.Response, category proto.Category, timePrefix string, page int) error {
	fileName := fmt.Sprintf("%s_%s_p%d.html", timePrefix, category.String(), page)
	fullPath := filepath.Join(task.outputDir, fileName)

	if err := r.Save(fullPath); err != nil {
		return err
	}
	log.Println("Saved", fullPath)
	return nil
}
