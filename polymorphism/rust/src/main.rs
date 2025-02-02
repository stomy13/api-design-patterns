use serde::{Deserialize, Serialize};
use serde_json;
use std::fmt::Debug;

fn main() {
    marshal_json();
    unmarshal_json();
}

fn marshal_json() {
    println!("### Marshal JSON");
    let chat_room = ChatRoom {
        messages: vec![
            Message::Text(TextMessage {
                id: "1".to_string(),
                message_type: "text".to_string(),
                text: "Hello, World!".to_string(),
            }),
            Message::Image(ImageMessage {
                id: "2".to_string(),
                message_type: "image".to_string(),
                image_uri: "https://example.com/image.jpg".to_string(),
            }),
            Message::Video(VideoMessage {
                id: "3".to_string(),
                message_type: "video".to_string(),
                video_uri: "https://example.com/video".to_string(),
            }),
        ],
    };

    let json = serde_json::to_string_pretty(&chat_room).unwrap();
    println!("{}", json);
}

fn unmarshal_json() {
    println!("### Unmarshal JSON");
    let json = r#"{
        "messages": [
            {
                "type": "Text",
                "id": "4",
                "message_type": "text",
                "text": "Hello, World!"
            },
            {
                "type": "Image",
                "id": "5",
                "message_type": "image",
                "image_uri": "https://example.com/image.jpg"
            },
            {
                "type": "Video",
                "id": "6",
                "message_type": "video",
                "video_uri": "https://example.com/video"
            }
        ]
    }"#;

    let chat_room: ChatRoom = serde_json::from_str(json).unwrap();
    println!("{:?}", chat_room);

    let json = serde_json::to_string_pretty(&chat_room).unwrap();
    println!("{}", json);
}

#[derive(Debug, Serialize, Deserialize)]
struct ChatRoom {
    messages: Vec<Message>,
}

#[derive(Debug, Serialize, Deserialize)]
#[serde(tag = "type")]
enum Message {
    Text(TextMessage),
    Image(ImageMessage),
    Video(VideoMessage),
}

#[derive(Debug, Serialize, Deserialize)]
struct TextMessage {
    id: String,
    message_type: String,
    text: String,
}

#[derive(Debug, Serialize, Deserialize)]
struct ImageMessage {
    id: String,
    message_type: String,
    image_uri: String,
}

#[derive(Debug, Serialize, Deserialize)]
struct VideoMessage {
    id: String,
    message_type: String,
    video_uri: String,
}
