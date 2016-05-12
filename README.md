# Apollo
Open source DJ for group situations

# Features
- Google Play Music
- Collaborative queue

# Infastructure
```

	    -------------------------
	   |  --------------- -----  |
	   | | Digital Ocean | AWS | |
	   |  --------------- -----  |
	   |                         |
	   | Host                    |
	    -------------------------
		      |
		      |
		      V
	    ---------------------
	   |  -------- --------  |
	   | | CoreOS | Ubuntu | |
	   |  -------- --------  |
	   |                     |
	   | OS                  |
	    ---------------------
		      |
		      |
		      V
                --------------
               | Apache Mesos |
                --------------
                      |
	              |
	              V
                 -------------
                | Kubernetes |
                 -------------
                      |
	              |
                      V
     ---------    ---------    --------- 
    | Service |  | Service |  | Service |
     ---------    ---------    ---------
```

# Behavior
- Master app will have complete control over queue
    - Remove/Add songs
    - Reorder songs
- Remove song from queue after a certain percentage of dislikes

# Extraneous Features
- Blacklist song
- Change dislike removal percentage

# Code
- Queue
  - String ms_id
    - Music source id
  - Song[] up_next
    - Next 3 songs to play, can not be changed
  - Song[] queue
    - General queue, most voted song gets popped off into up_next
  - int up_next_size
    - Size of the up next queue
- Song
  - String ms_id

  - String id
  - String name

  - String album_id
  - String album_name

  - String artwork_url
  - double length_ms

  - int likes
  - int dislikes
- MusicSource
  - Song[] query(String query)
  - Song get_song(String id)
  - String get_song_streaming_url(String id, StreamQuality quality)
- StreamQuality
  - LOW
  - MED
  - HIGH
