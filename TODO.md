
# Send daily emails

## Start with lambda for Moby Dick writing to s3

* make pick and check s3 instead of disk DONE
* read from s3 (print to be sure) DONE
* write back to s3 DONE
* put this in Lambda on cron

## Eventual Plan

* store emails in s3
* each time it runs
    * make pick and get text from s3
    * query db/s3 for emails
    * send pick to all emails (instead of writing to s3)
        * DO I NEED SES FOR THIS???

# Improvements and Extensions

## Multiple books

Make a lambda that will call `GetPickS3()` for each book and send it to all subscribers.

## Sign up

* a way to see what books are available
* sign up for a given book (send in your email)
* ideally a way to request a book (see below)


## Get Books

* getbook() should write to an s3 location
* a way to trigger get book (from an API call?)
  * we may not care that much at this point. I can just do it locally.

## Data Storage, etc.

* store emails in RDS? Probably not. But maybe something more interesting than just csv/json in s3.
