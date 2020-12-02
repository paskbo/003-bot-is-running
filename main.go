package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const EpicTokenChvfefe string = "Nzc4NzEwMzgyNjEyMDU0MDE2.X7V8jg.lfmSHRXWybSDmH4fbd1VY2j5Xlw"

var BotID string

/*What we talked the other day.
Since BotID it starts with a capital B, it's gonna be a global thing.
I've looked through some examples from discordgo,
in order to understand the fundamental of the following:*/

func main() {
	dg, err := discordgo.New("Bot " + EpicTokenChvfefe) //I honestly want to make a config file for that

	if err != nil { //Error log for the cmd, if the token is invalid. In this case, it will be.
		fmt.Println(err.Error())
		return //To close, makes no sense keeping it running.
	}

	u, err := dg.User("@me") //In order to get the Bot ID

	if err != nil { //Error log for cmd
		fmt.Println(err.Error())
	}

	BotID = u.ID

	dg.AddHandler(chvfefeHandler)
	err = dg.Open() //Should open the connection

	if err != nil { //Error log for cmd connecting
		fmt.Println(err.Error())
		return // If it doesn't connect, it should logically exit.
	}

	fmt.Println("A text message that chvfefe is running, for example. Was it chvfefe?")
	//Like in the JavaScript Bot I had, tells you the bot is working if no error encountered

	<-make(chan struct{})
	return
	/*I had no idea about line 39, I've read that recently.
	  They say this keeps the bot running Like a handler I suppose?*/

	/* I had a bad week, so I'm a little bit late. Right now as you may read this,
	   I might be working on this now. That's just a test bot, still looking on how to send messages.*/
}

func chvfefeHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(m.Content)
}
