## What
This is a terminal tool for downloading drum maps and installing them to your DTXManiaNX files folder.
## Why
There is a useful DTXMania drum maps website https://approvedtx.blogspot.com by @approvedtx . You can get the hottest and latest songs' DTXMania maps(all excellent) from it. 

However, it is trivial to download them one by one from Google Drive. You need to click the link first, redirect to another page and click confirm button most of the time. \
After downloading them, you also need to unzip and put them to DTXMania's files folder manually.

Now you can do all these things by using only a few commands.
## How to use
1. You should set your DTXManiaNX files folder's path to the tool by two ways:
   1. Use  ``DTXMapDownloader_windows-amd64.exe config game "<files folder path>"`` to set it.
   2. Open ``setting.json`` file(it will be created after running the tool first time) and change the value of ``game_songs_path`` to your ``<files folder path>``.
2. You can use ``DTXMapDownloader_windows-amd64.exe search "<song name>"`` to search the song drum map you want to get.
3. Then you can use ``DTXMapDownloader_windows-amd64.exe download "<song name>"`` to download and install drum map. \
Actually, the tool will search it and download the first song map from the results.
4. Run the game, and you can see the new song.
## Thanks
Thanks to @approvedtx. Please support him more!