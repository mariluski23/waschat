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