import json
import os
import shutil

def f1():
    dirs = [x for x in os.listdir('.') if os.path.isdir(x)]

    for name in dirs:
        mod = '_' + name
        os.rename(name, mod)
        revs = os.listdir(os.path.join(mod, 'revs'))
        highest = sorted(revs)[-1]
        os.rename(os.path.join(mod, 'revs', highest), name)

def f2():
    for fn in os.listdir('.'):
        if os.path.isdir(fn):
            continue

        print(fn)
        with open(fn, 'r+') as f:
            try:
                data = json.load(f)
            except ValueError:
                continue

            f.seek(0)
            f.truncate()
            data = data['Content'].replace("\r\n", "\n")
            f.write(data)

f2()
