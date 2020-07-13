
# Send daily emails

Make a lambda that will call `GetPick()` for each book and send it to all subscribers.

Or should there be a separate lambda for each book? Maybe that's the skateboard.

* store emails in s3 or RDS
* each time it runs
    * query s3 for book
    * `GetPick()` on that book
    * query db for emails associated with that book
    * send pick to all emails
        * DO I NEED SES FOR THIS???
        * maybe start with reading and writing the pick back to somewhere else in s3, just to test
    * go to next book (maybe)

# Sign up

* a way to see what books are available
* sign up for a given book (send in your email)
* ideally a way to request a book (see below)


# Get Books

* getbook() should write to an s3 location
* a way to trigger get book (from an API call?)
  * we may not care that much at this point. I can just do it locally.