/*
                       WasChat
                       ========

A FOSS, decentrailzed, terminal chat client using WebSockets.
Written in Go.

Under the GNU GPLv3 license.
*/

package main

// Import necessary packages
import (
    // I/o
    "fmt"
    // OS things
    "os"
    // WebSocket
    "github.com/gorilla/websocket"
    // BubbbleTea (terminal UI) and LipGloss (styling)
    "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
    // Networking
    "net"
    "net/http"
    // UI Components
    "github.com/mariluski23/waschat/components"
)

// Test
func main() {
    // Create a new bubbletea instance
    t, err := bubbletea.New()
    if err != nil {
        fmt.Println("Error creating bubbletea instance:", err)
        os.Exit(1)
    }

    // Create a new lipgloss context
    ctx := lipgloss.NewContext()

    // Create a new WebSocket connection
    ws, err := websocket.Dial("ws://localhost:8080/ws", "", "http://localhost")
    if err != nil {
        fmt.Println("Error creating WebSocket connection:", err)
        os.Exit(1)
    }
    defer ws.Close()

    // Create a new bubbletea terminal
    term := t.NewTerminal(ws, "main")

    // Create a new chat window
    chat := components.NewChatWindow(term, ctx)

    // Create a new input window
    input := components.NewInputWindow(term, ctx)

    // Create a new button component
    button := components.NewButton("Send", [3]int{0, 255, 0}, func() {
        // Send the message
        ws.WriteMessage(websocket.TextMessage, []byte(input.GetInput()))
        // Clear the input
        input.ClearInput()
    }, 10, 1)

    // Create a new layout
    layout := components.NewLayout(chat, input, button)

    // Start the bubbletea terminal
    if err := term.Start(); err != nil {
        fmt.Println("Error starting bubbletea terminal:", err)
        os.Exit(1)
    }

    // Start the bubbletea event loop
    if err := t.Start(); err != nil {
        fmt.Println("Error starting bubbletea event loop:", err)
        os.Exit(1)
    }

    // Start the bubbletea layout
    if err := layout.Start(); err != nil {
        fmt.Println("Error starting bubbletea layout:", err)
        os.Exit(1)
    }
}