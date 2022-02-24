# mapimage [![CircleCI](https://circleci.com/gh/danesparza/mapimage.svg?style=shield)](https://circleci.com/gh/danesparza/mapimage)
AWS Lambda handler to get a map image for a set of coordinates.  

Encodes as a data uri (with image/jpeg) because returning image data on one handler in API gateway [seems to require settings changes that then require changes on other resources on the same API](https://stackoverflow.com/a/50670252/19020).  Returning as base64 encoded JSON data.
