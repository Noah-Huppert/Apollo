from gmusicapi import Mobileclient

api = Mobileclient()
api.login("noahhuppert@gmail.com", "osilgdgrnmiigtfz", Mobileclient.FROM_MAC_ADDRESS)

print(api.search_all_access("Hello"))
