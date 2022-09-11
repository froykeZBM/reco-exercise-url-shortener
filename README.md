# reco-exercise-url-shortener
Backend Exercise for Reco
The plan was to divide the program to the following modules:
- Storage: handle URL DB
  - Right now it is a map.
  - Later it could be changed to a json file / REDIS DB
  - Additional issues: 
    - colliding IDS 
    - Detecting Long URL that has been already posted
- url generator - generating short url out of long ones
  - Implementation - a hash function
- redirecting - getting the original URL given the short one
- handling http function 
  - get, post, and redirecting the package

