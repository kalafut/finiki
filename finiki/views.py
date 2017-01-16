import os
from contextlib import contextmanager
from finiki import app
from flask import redirect, render_template, request

import mistune

markdown = mistune.Markdown()

ROOT = '/Users/kalafut/Dropbox/finiki'
DEFAULT_EXT = 'md'
RECENT_CNT = 7

@app.route('/', defaults={'path': ''})
@app.route('/<path:path>', methods=['GET', 'POST'])
def index(path):
    isdir = os.path.isdir(tod(path))
    if isdir and not path.endswith('/') and path != '':
        return redirect(path + '/'), 303

    if request.method == 'POST':
        if 'delete' in request.form:
            os.remove(tof(path))
            return redirect('/')
        else:
            with opener(path, 'w') as f:
                f.write(request.form['text'])
                return redirect(path)

    if isdir:
        d, p = scan(path)
        return render_template('dir.html', dirs=d, pages=p)
    action = request.args.get('action')

    if action == 'edit':
        with opener(path) as f:
            contents = f.read()
            return render_template('edit.html', text=contents, path=path)
    elif action == 'delete':
        return render_template('delete.html', path=path)

    try:
        with opener(path) as f:
            contents = f.read()
            return render_template('show.html', text=markdown(contents), path=path)
    except NotADirectoryError:
        msg = "You cannot have paths under a page."
        return render_template('error.html', message=msg)
    except FileNotFoundError:
        contents = 'New Page'
        return render_template('edit.html', text=contents, path=path)


def scan(path):
    d, p = [], []
    for entry in os.scandir(tod(path)):
        if not entry.name.startswith(('.', '__')):
            if entry.is_dir():
                d.append(entry.name)
            else:
                p.append(os.path.splitext(entry.name)[0])
    return d, p


def load_recent(skip_first=False, recent_cnt=RECENT_CNT):
    with open(tof('__system/recent')) as f:
        lines = f.readlines()
        start = 1 if skip_first else 0
        return (x.strip() for x in lines[start:start + recent_cnt])


@contextmanager
def opener(path, mode='r'):
    if mode == 'w':
        os.makedirs(os.path.dirname(tof(path)), exist_ok=True)
    with open(tof(path), mode) as f:
        yield f

def tof(path):
    return "{}.{}".format(os.path.join(ROOT, path), DEFAULT_EXT)

def tod(path):
    return os.path.join(ROOT, path)

