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
                | Kubernertes |
                 -------------
                      |
	              |
                      V
     ---------    ---------    --------- 
    | Service |  | Service |  | Service |
     ---------    ---------    ---------
```

# Architecture
- DJ
  - Google Play Music Source

# Code
- Queue
  - String ms_id
  - Song[] up_next
    - Next 3 songs to play, can not be changed
  - Song[] queue
    - General queue, most voted song gets popped off into up_next
- Song
  - String ms_id

  - String id
  - String name

  - String album_id
  - String album_name

  - String artist_id
  - String artist_name

  - String artwork_url
  - double length_ms

  - int votes
- MusicSource
  - Song[] query(String query)
  - Song get_song(String id)
  - String get_song_streaming_url(String id, StreamQuality quality)
- StreamQuality
  - LOW
  - MED
  - HIGH
