import { useState } from "react";
import {
  Box,
  TextField,
  Button,
  Typography,
  Paper,
  List,
  ListItem,
  Divider,
} from "@mui/material";

export default function Home() {
  const mockMessages = [
    { username: "Alice", message: "Hey there!" },
    { username: "Bob", message: "Hi Alice, how are you?" },
    { username: "Alice", message: "I'm good, just trying out this chat app." },
    { username: "Charlie", message: "Hello everyone!" },
  ];

  const [messages, setMessages] = useState(mockMessages);
  const [messageInput, setMessageInput] = useState("");
  const [username, setUsername] = useState("You");

  const handleSendMessage = () => {
    if (messageInput.trim()) {
      const newMessage = { username, message: messageInput };
      setMessages((prev) => [...prev, newMessage]);
      setMessageInput("");
    }
  };

  return (
    <Box
      sx={{
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "space-between",
        minHeight: "100vh",
        p: 4,
      }}
    >
      <Box
        sx={{
          width: "100%",
          maxWidth: "800px",
          textAlign: "center",
        }}
      >
        <Typography variant="h4" fontWeight="bold" gutterBottom>
          Go Chat Room
        </Typography>
        <Paper
          elevation={3}
          sx={{
            width: "100%",
            p: 2,
            maxHeight: "400px",
            overflowY: "auto",
            backgroundColor: "#f5f5f5",
            mb: 2,
          }}
        >
          <List>
            {messages.map((msg, index) => (
              <div key={index}>
                <ListItem
                  sx={{
                    p: 1,
                    mb: 1,
                    backgroundColor: "white",
                    borderRadius: "8px",
                  }}
                >
                  <Typography
                    variant="subtitle2"
                    sx={{ fontWeight: "bold", color: "blue" }}
                  >
                    {msg.username}:
                  </Typography>
                  <Typography variant="body2" sx={{ ml: 1 }}>
                    {msg.message}
                  </Typography>
                </ListItem>
                {index < messages.length - 1 && <Divider />}
              </div>
            ))}
          </List>
        </Paper>

        <Box
          sx={{
            display: "flex",
            gap: 1,
          }}
        >
          <TextField
            fullWidth
            value={messageInput}
            onChange={(e) => setMessageInput(e.target.value)}
            onKeyPress={(e) => e.key === "Enter" && handleSendMessage()}
            placeholder="Type a message..."
            variant="outlined"
            size="small"
          />
          <Button
            variant="contained"
            color="primary"
            onClick={handleSendMessage}
          >
            Send
          </Button>
        </Box>
      </Box>
    </Box>
  );
}
