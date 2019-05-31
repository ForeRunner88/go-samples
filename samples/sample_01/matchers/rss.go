package matchers

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/ForeRunner88/go-samples/samples/sample_01/search"
)

type (
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubData     string   `xml:"pubData"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubData        string   `xml:"pubData"`
		LastBuildData  string   `xml:"lastBuildData"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:channel`
	}
)

type rssMatcher struct{}

func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

func (m rssMatcher) retrive(feed *search.Feed) (*rssDocument, error) {
	if feed.URL == "" {
		return nil, fmt.Errorf("No rss feed URL")
	}
	resp, err := http.Get(feed.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response ERROR %d\n", resp.StatusCode)
	}

	var doc rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&doc)
	return &doc, err
}

func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result
	log.Printf("Search Feed Type[%s] Site[%s] For URL[%s]\n",
		feed.Type, feed.Name, feed.URL)
	doc, err := m.retrive(feed)
	if err != nil {
		return nil, err
	}
	for _, channelitem := range doc.Channel.Item {
		matched, err := regexp.MatchString(searchTerm, channelitem.Title)
		if err != nil {
			return nil, err
		}
		if matched {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelitem.Title,
			})
		}
		matched, err = regexp.MatchString(searchTerm, channelitem.Description)
		if err != nil {
			return nil, err
		}
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelitem.Description,
			})
		}
	}
	return results, nil
}
