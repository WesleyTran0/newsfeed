# Newsfeed - A CyberNews Web Scraper

Newsfeed is a web scraper written entirely in go. Staying up to date in cybersecurity is critical, but subscribing to
multiple newsletters can lead to email clogging. This project instead scrapes for articles directly from the urls,
providing headlines, dates, and descriptions in a JSON format to be exported. This project is not intended to provide an
easy way to view these articles. Instead, it is meant to coalesce articles from multiple sites in an easily parsed
format, so that it articles can be displayed elsewhere.

### Future Implementations

- Rather than exporting as a json file, I want to eventually have it update a database and possibly link the database to some web service that I own
- In addition to all the information I gather, I want to include the pictures of thumbnails to allow for a better view 
