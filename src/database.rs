use anyhow::Ok;
use anyhow::Result;
use libsql_client::Statement;
use libsql_client::{client::GenericClient, DatabaseClient};

use crate::smtp::Mail;

pub struct Client {
    db: GenericClient,
}

impl Client {
    pub async fn new() -> Result<Self> {
        let mut path = std::env::temp_dir();
        path.push("haxmail.db");

        let db = libsql_client::new_client().await?;
        db.batch([
            "CREATE TABLE IF NOT EXISTS mail (date text, sender text, rcpt text,data text)",
            "CREATE INDEX IF NOT EXISTS mail_date ON mail(date)",
            "CREATE INDEX IF NOT EXISTS mail_rcpt ON mail(rcpt)",
        ])
        .await?;
        Ok(Self { db })
    }

    pub async fn replicate(&self, mail: Mail) -> Result<()> {
        let now = chrono::offset::Utc::now()
            .format("%Y-%m-%d %H:%M:%S%.3f")
            .to_string();

        self.db
            .execute(Statement::with_args(
                "INSERT INTO mail VALUES (?,?,?,?)",
                libsql_client::args!(now, mail.from, mail.to.join(", "), mail.data),
            ))
            .await
            .map(|_| ())
    }
}
