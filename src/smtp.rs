use std::io::{BufRead, Error, Write};

//client commands
const HELO_START: &str = "HELO ";
const MAIL_START: &str = "MAIL FROM:";
const RCPT_START: &str = "RCPT TO:";
const DATA_LINE: &str = "DATA";
const QUIT_LINE: &str = "QUIT";

//Server responses
const MSG_READY: &str = "220 ready";
const MSG_OK: &str = "250 OK";
const MSG_SEND_MESSAGE_CONTENT:&str="354 Send message content";
const MSG_BYE:&str="221 Bye";
const MSG_SYNTAX_ERROR:&str="500 unexpected line";

pub struct Message {
    sender: String,
    receiver: Vec<String>,
    data: Vec<String>,
}

enum State {
    Helo,
    Mail,
    Rcpt,
    RcptOrData,
    Dot,
    MailOrQuit,
    Done,
}

pub struct Connection {
    state: State,
    sender_domain: String,
    messages: Vec<Message>,
    next_sender: String,
    next_recipients: Vec<String>,
    next_data: Vec<String>,
}


impl Connection {
    pub fn new() -> Connection {
        Connection {
            state: State::Helo,
            sender_domain: "".to_string(),
            messages: Vec::new(),
            next_sender: "".to_string(),
            next_recipients: Vec::new(),
            next_data: Vec::new(),
        }
    }
    fn feed_line<'a>(&mut self,line: &'a str)->Result<&'a str, &'a str>{
            match self.state {
                State::Helo=>{
                    if line.starts_with(HELO_START){
                        self.sender_domain=line[HELO_START.len()..].trim().to_string();
                        self.state=State::Mail;
                        Ok(MSG_OK)
                    }
                    else{
                        Err(MSG_SYNTAX_ERROR)
                    }
                }
            }
    }
    pub fn handle(reader: &mut dyn BufRead, writer: &mut dyn Write) -> Result<Connection, Error> {
        let mut result = Connection::new();

        writeln!(writer, "{}", MSG_READY);
        loop {
            let mut line = String::new();
            reader.read_line(&mut line);
            match result
        }
        Ok(result)
    }
}
