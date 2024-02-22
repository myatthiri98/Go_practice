package main

import (
	"log"
	"net/http"
)

// func main() {

// 	// name := "Go"

// 	// var someName string = "Java"

// 	// fruits := []string{"apple","orange"}
// 	// fmt.Println("Hello World")

// 	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	w.Write([]byte("Hello from Myanmar"))
// 	// })

// 	// http.Handle("/", http.FileServer(http.Dir("templates")))

// 	// err := http.ListenAndServe(":8000", nil)
// 	// if err != nil {
// 	// 	log.Fatal("ListenAndServe:", err.Error())
// 	// }

// if err := http.ListenAndServe(":8000", nil); err != nil {
// 	log.Fatal("ListenAndServe:", err.Error())
// }
// }

type canItBeHandler struct {
}

func (h *canItBeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Simple Chatbot</title>
    <style>
	body {
		font-family: Arial, sans-serif;
		padding: 100px;
		display: flex;
		justify-content: center;
		align-items: center;
		font-size: 32px,
	}
	
	.chat-container {
		max-width: 400px;
		margin: auto;
		border: 1px solid #ccc;
		border-radius: 5px;
		padding: 10px;
	}
	
	.chat-box {
		background-color: #f1f1f1;
		margin-bottom: 10px;
		padding: 10px;
		border-radius: 5px;
	}
	
	.user {
		color: #0a0;
	}
	
	.bot {
		color: #00a;
	}
	
	.input-box {
		width: calc(100% - 22px); /* Adjusted width to accommodate the button margin */
		padding: 5px;
	}
	
	.btn {
		margin-top: 10px;
		width: 100%;
		padding: 10px;
		background-color: #0a0;
		color: white;
		border: none;
		border-radius: 5px;
		cursor: pointer;
	}	
    </style>
</head>
<body>
    <div class="chat-container">
        <div class="chat-box bot">Hi there! I'm a simple chatbot. How can I help you today?</div>
        <input type="text" class="input-box" id="userInput" placeholder="Type your message here...">
        <button class="btn" onclick="sendMessage()">Send</button>
    </div>

    <script>
        function sendMessage() {
            var userInput = document.getElementById("userInput").value;
            var chatBox = document.createElement("div");
            chatBox.textContent = "You: " + userInput;
            chatBox.className = "chat-box user";
            document.getElementById("userInput").value = "";
            document.querySelector(".chat-container").appendChild(chatBox);
            getBotResponse(userInput);
        }

        function getBotResponse(userInput) {
            var botResponse = "I'm sorry, I don't understand. Can you rephrase your question?";
            if (userInput.includes("hello")) {
                botResponse = "Hello! How can I assist you today?";
            } else if (userInput.includes("how are you")) {
                botResponse = "I'm just a bot, so I don't have feelings. But thank you for asking!";
            } else if (userInput.includes("bye")) {
                botResponse = "Goodbye! Have a great day!";
            }
            var chatBox = document.createElement("div");
            chatBox.textContent = "Bot: " + botResponse;
            chatBox.className = "chat-box bot";
            document.querySelector(".chat-container").appendChild(chatBox);
        }
    </script>
</body>
</html>
`))
}

func main() {
	http.Handle("/", &canItBeHandler{})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}
