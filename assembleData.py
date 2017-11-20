from mutagen.mp3 import MP3
import glob, json
data = []
for name in glob.glob('music/*.mp3'):
    audio = MP3("./" + name)
    data.append({
                'length':"{0}m {1}s".format(int(audio.info.length // 60), int(audio.info.length % 60)), 
                'name':name.split("music\\")[1]
                })
    print(data[-1])
    #print(name.split("music\\")[1])
    #audio = MP3("./" + name)
    ##print("{0}m {1}s".format(int(audio.info.length // 60), int(audio.info.length % 60))
with open("data.json", "w") as of:
    json.dump(data, of)