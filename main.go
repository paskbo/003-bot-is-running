package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const EpicTokenChvfefe string = "Did you know strawberries aren't berries?"

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
	if m.Author.ID == BotID { //If the ID of the Author of the Message is equal to the Bot ID that is set
		return
		/*To not execute the rest of the code, just in case it exits.
		So the Bot doesn't talk to itself, just want to avoid any spam problems here*/
	}

	if m.Content == "chvfefe" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Yep, you're definitely right") //Change m.ChannelID to "Here-Your-Channel-ID", if you only want it to work on one channel. Good for a future application bot,
	/* Yes, I know. I used the blank identifier, in order to avoid any "undeclared" variable error, happened to me already.
	I found the source here:
	https://www.geeksforgeeks.org/what-is-blank-identifierunderscore-in-golang/#:~:text=_(underscore)%20in%20Golang%20is,unused%20variable%20using%20Blank%20Identifier.&text=It%20hides%20the%20variable%27s%20values%20and%20makes%20the%20program%20readable.*/
	}
}

