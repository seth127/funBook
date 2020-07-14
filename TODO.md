
# Send daily emails

## Start with lambda for Moby Dick writing to s3

* make pick and check s3 instead of disk
* read from s3 (print to be sure)
* write back to s3 DONE
* put this in Lambda on cron

## Eventual Plan

* store emails in s3
* each time it runs
    * make pick and check s3 for text
    * query db for emails
    * send pick to all emails (instead of writing to s3)
        * DO I NEED SES FOR THIS???

## Improvements

* store emails in RDS?

## Multiple books

Make a lambda that will call `GetPick()` for each book and send it to all subscribers.

# Sign up

* a way to see what books are available
* sign up for a given book (send in your email)
* ideally a way to request a book (see below)


# Get Books

* getbook() should write to an s3 location
* a way to trigger get book (from an API call?)
  * we may not care that much at this point. I can just do it locally.