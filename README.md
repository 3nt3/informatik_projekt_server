# informatik_projekt_server

CS Project made by [me](https://niels-dingsbums.de) and [Sheesher](http://ichbindumm12321.de).

### Use
It is a game-server-api built for minigames (tic tac toe, rock paper scisors, [etc](https://github.com/NielsDingsbums/informatik_projekt_server/branches)).

---

#### Example output
```shell
2018/11/06 15:06:08 Added (rps) room: {0 [{0 niels 0 -1} {1 someone_else 0 -1}]}
2018/11/06 15:06:16 Updated fig of player "niels" to 1
2018/11/06 15:06:19 GET figs of room 0 [{0 niels 0 1} {1 someone_else 0 -1}] -> [1 -1]
```


### Usage
```shell
$ go build
$ sudo ./informatik_projekt_server
```

