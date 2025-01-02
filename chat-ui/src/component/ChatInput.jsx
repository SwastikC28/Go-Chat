import React, { useState } from "react";
import { Container, Row, Col, Form, Button, Card } from "react-bootstrap";
import { connect,sendMsg } from '../api/index.js'
import { useEffect } from 'react';

const ChatApp = () => {
    const [messages, setMessages] = useState([]);
    const [input, setInput] = useState("");

    // Handle sending a message
    const handleSendMessage = () => {
        if (input.trim() === "") return;
        const newMessage = {
            text: input,
            timestamp: new Date().toLocaleTimeString(),
            isMine: true, // Mark this message as sent by the user
        };
        setMessages([...messages, newMessage]);
        sendMsg(newMessage)
        setInput("");
    };

    // Handle Enter key press
    const handleKeyDown = (e) => {
        if (e.key === "Enter") {
            e.preventDefault(); // Prevent default Enter behavior (e.g., form submission)
            handleSendMessage();
        }
    };

    const handler = (msg) => {
        console.log("Pokemon", msg);
    }


    useEffect(() => {
        connect(handler)
    }, [])

    return (
        <Container className="mt-4" style={{ maxWidth: "55vw" }}>
            <Card>
                <Card.Header as="h5">Chat Application</Card.Header>
                <Card.Body>
                    {/* Message Display Area */}
                    <div
                        style={{
                            height: "60vh", // 60% of viewport height
                            minHeight: "400px",
                            overflowY: "auto",
                            border: "1px solid #ddd",
                            borderRadius: "4px",
                            padding: "10px",
                            backgroundColor: "#f8f9fa",
                        }}
                    >
                        {messages.length === 0 ? (
                            <p className="text-muted">No messages yet. Start the conversation!</p>
                        ) : (
                            messages.map((msg, index) => (
                                <div
                                    key={index}
                                    className={`d-flex ${msg.isMine ? "justify-content-end" : "justify-content-start"} mb-2`}
                                >
                                    <div
                                        style={{
                                            maxWidth: "75%",
                                            padding: "10px",
                                            borderRadius: "10px",
                                            backgroundColor: msg.isMine ? "#007BFF" : "#E9ECEF",
                                            color: msg.isMine ? "white" : "black",
                                        }}
                                    >
                                        <small className="d-block text-muted">{msg.timestamp}</small>
                                        <p className="mb-0">{msg.text}</p>
                                    </div>
                                </div>
                            ))
                        )}
                    </div>

                    {/* Input and Send Area */}
                    <Row className="mt-3">
                        <Col xs={9}>
                            <Form.Control
                                type="text"
                                placeholder="Type a message..."
                                value={input}
                                onChange={(e) => setInput(e.target.value)}
                                onKeyDown={handleKeyDown} // Attach keydown event here
                            />
                        </Col>
                        <Col xs={3}>
                            <Button variant="primary" onClick={handleSendMessage} className="w-100">
                                Send
                            </Button>
                        </Col>
                    </Row>
                </Card.Body>
            </Card>
        </Container>
    );
};

export default ChatApp;
