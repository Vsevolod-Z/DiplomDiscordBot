package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"discordbot/models"

	"github.com/bwmarrin/discordgo"
)

var dg *discordgo.Session

func main() {
	var err error
	dg, err = discordgo.New("Bot token")
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}
	dg.AddHandler(handleMessage)
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
		return
	}
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	if len(m.Mentions) > 0 && m.Mentions[0].ID == s.State.User.ID {
		s.ChannelMessageSend(m.ChannelID, "hi")
		return
	}
	args := strings.Split(m.Content, " ")
	if args[0] == "/title" {
		if len(args) < 2 {
			s.ChannelMessageSend(m.ChannelID, "Please specify a title.")
			return
		}
		title := strings.Join(args[1:], " ")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Title: %s", title))
	} else if args[0] == "/ids" {
		if len(args) < 2 {
			s.ChannelMessageSend(m.ChannelID, "Please specify an ID.")
			return
		}
		id := strings.Join(args[1:], " ")
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("ID: %s", id))
	}
}
func isBotMentioned(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	for _, user := range m.Mentions {
		if user.ID == s.State.User.ID {
			return true
		}
	}
	return false
}
func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if !isBotMentioned(s, m) {
		return
	}
	fmt.Println(m.Content)
	if m.Content == "Hi" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hi Back")
	}
	fmt.Println(m.Content)
	parts := strings.Split(m.Content, " ")
	fmt.Println(parts)
	fmt.Println(parts[0])
	fmt.Println(parts[1])
	fmt.Println("cmd", parts[1])
	fmt.Println(parts[2])
	if len(parts) < 2 || !strings.HasPrefix(parts[1], "/") {
		return
	}
	cmd := strings.ToLower(parts[1])
	fmt.Println(cmd)
	switch cmd {
	case "/title":
		count := 0
		fmt.Println(parts[2:])
		str := strings.Join(parts[2:], " ")
		for _, data := range FindByTitle(str) {
			SendUserAppData(data, m.ChannelID)
			count++
			if count == 10 {
				break
			}
		}
	case "/ids":
		gameID := parts[2]
		fmt.Println(gameID)
		for _, data := range FindById(gameID) {
			SendUserAppData(data, m.ChannelID)
		}
	default:
		_, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("шото не так"))
		if err != nil {
			log.Println("error sending message: ", err)
		}
		return
	}
}

func splitCommandAndArgument(text string) (string, string) {
	spaceIndex := strings.Index(text, " ")
	if spaceIndex == -1 {
		return text[:spaceIndex+1], text[spaceIndex+1:]
	}
	fmt.Println(text[:spaceIndex+1], text[spaceIndex+1:])
	return text[:spaceIndex], text[spaceIndex+1:]
}

func SendUserAppData(data models.AppMainData, channelID string) {
	response, err := http.Get(data.SteamAppData.Icon)
	if err != nil {
	}
	defer response.Body.Close()
	tmpFile, err := os.CreateTemp("", "image_*.jpg")
	if err != nil {
	}
	defer os.Remove(tmpFile.Name())
	_, err = io.Copy(tmpFile, response.Body)
	if err != nil {
	}
	gogPrice, err := strconv.ParseFloat(data.GOGAppData.Price.FinalAmount, 64)
	steamBuyPrice, err := strconv.Atoi(data.SteamBuyAppData.Price.Rub)
	rating := int((float64(data.SteamAppData.AppReview.QuerySummary.TotalPositive) / float64(data.SteamAppData.AppReview.QuerySummary.TotalReviews)) * 100)
	var windows string
	var mac string
	var linux string
	if data.SteamAppData.Platforms.Windows {
		windows = "✅"
	} else {
		windows = "❌"
	}
	if data.SteamAppData.Platforms.Mac {
		mac = "✅"
	} else {
		mac = "❌"
	}
	if data.SteamAppData.Platforms.Linux {
		linux = "✅"
	} else {
		linux = "❌"
	}
	caption := fmt.Sprintf("Название: %s\n\n Рейтинг игры в Steam: %d%% \n\nWindows:%s\nMac:%s\nLinux:%s\n\n Цены,которые мы нашли:\n Steam: %s \n GOG: %s \n SteamPAY: %s \n SteamBUY: %s", data.SteamAppData.Name, rating, windows, mac, linux, addLink(data.SteamAppData.URL, data.SteamAppData.SteamPriceOverview.GameFinalPrice), addLink(data.GOGAppData.URL, int(math.Round(gogPrice))), addLink(data.SteamPayAppData.URL, data.SteamPayAppData.Prices.Rub), addLink(data.SteamBuyAppData.URL, steamBuyPrice))
	dg.ChannelMessageSend(channelID, caption)
}
func addLink(url string, price int) string {
	word := strconv.Itoa(price)
	return fmt.Sprintf("%sр \nURL:%s", word, url)
}
func FindByTitle(title string) []models.AppMainData {
	fmt.Println(title)
	url := fmt.Sprintf("http://109.254.9.58:8080/api/apps/findByTitle?title=%s", strings.Replace(title, " ", "%20", -1))
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to send request: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var apps []models.AppMainData
	err = json.Unmarshal([]byte(body), &apps)
	if err != nil {
		log.Printf("")
	}
	return apps
}
func FindById(input string) []models.AppMainData {
	fmt.Println(input)
	url := fmt.Sprintf("http://109.254.9.58:8080/api/apps/findByIds?appids=%s", input)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Failed to send request: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var apps []models.AppMainData
	err = json.Unmarshal([]byte(body), &apps)
	if err != nil {
		log.Printf("")
	}
	return apps
}
