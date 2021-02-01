## **Media Server Automations**

This is a simple web application that hosts a set of automations which help in my home media server setup.

The automations include:
1. ##### Manage Plex media library by creating (or deleting) symbolic links to the actual media files.
    Plex server requires you to manage your media files in a certain library tree and by using a naming convention. See [this for Movies](https://support.plex.tv/articles/naming-and-organizing-your-movie-media-files/) and [this for TV Shows](https://support.plex.tv/articles/naming-and-organizing-your-tv-show-files/), and often that not, your media files are stored randomly due to various reasons (torrent downloads, space constraints, etc.)
    
    _Media Server Automations_ solves this by giving you a UI to manage your plex library (Movies and TV shows) by offering a simple UI.
   
    The UI scans a set of directory and offers you to either add any directory or file as a Movie or a TV Show (with seasons) by cleverly creating symlinks to the actual files.
    
    The UI also offers you to delete any of the titles from the library and keeping the underlying files and deleting the structure it created for the plex library
   