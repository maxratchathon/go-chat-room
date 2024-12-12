import { useState } from 'react';
import { useWebSocket } from '../src/components/WebSocketProvider';

export default function Home() {
  const { messages, sendMessage } = useWebSocket();
  const [messageInput, setMessageInput] = useState('');

  const handleSendMessage = () => {
    if (messageInput.trim()) {
      sendMessage(messageInput);
      setMessageInput('');
    }
  };

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div className="z-10 max-w-5xl w-full items-center justify-between font-mono text-sm lg:flex">
        <h1 className="text-4xl font-bold mb-6">Go Chat Room</h1>
        
        <div className="message-container bg-gray-100 w-full p-4 rounded-lg h-96 overflow-y-auto">
          {messages.map((msg, index) => (
            <div key={index} className="message mb-2 p-2 bg-white rounded">
              {msg}
            </div>
          ))}
        </div>

        <div className="input-container mt-4 flex">
          <input 
            type="text"
            value={messageInput}
            onChange={(e) => setMessageInput(e.target.value)}
            onKeyPress={(e) => e.key === 'Enter' && handleSendMessage()}
            className="flex-grow p-2 border rounded-l"
            placeholder="Type a message..."
          />
          <button 
            onClick={handleSendMessage}
            className="bg-blue-500 text-white p-2 rounded-r"
          >
            Send
          </button>
        </div>
      </div>
    </main>
  );
}