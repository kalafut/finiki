#!/usr/bin/env python3

import os
from collections import OrderedDict
from contextlib import contextmanager
from flask import Flask, redirect, render_template, request
import jinja2

import mistune

app = Flask(__name__)
markdown = mistune.Markdown(hard_wrap=True)

try:
    ROOT = os.environ['FINIKI_ROOT']
except KeyError:
    print('Set the FINIKI_ROOT environment variable to your document root.')
    exit(1)

DEFAULT_EXT = 'md'
RECENT_CNT = 8


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
        return render_template('dir.html', dirs=d, pages=p, recents=load_recent(skip_first=False))

    action = request.args.get('action')

    if action == 'edit':
        with opener(path) as f:
            contents = f.read()
            return render_template('edit.html', text=contents, path=path, title=os.path.basename(path))
    elif action == 'delete':
        return render_template('delete.html', path=path)

    try:
        with opener(path) as f:
            contents = f.read()
            save_recent(path)
            return render_template('show.html', text=markdown(contents), path=path, recents=load_recent(skip_first=True), title=os.path.basename(path))
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
        return [x.strip() for x in lines[start:start + recent_cnt]]


def save_recent(path):
    recents = OrderedDict.fromkeys([path] + load_recent())
    with open(tof('__system/recent'), 'w') as f:
        f.write('\n'.join(recents.keys()))


@contextmanager
def opener(path, mode='r'):
    if mode == 'w':
        os.makedirs(os.path.dirname(tof(path)), exist_ok=True)
    with open(tof(path), mode) as f:
        yield f


@app.template_filter('basename')
def reverse_filter(s):
    return os.path.basename(s)


def tof(path):
    return "{}.{}".format(os.path.join(ROOT, path), DEFAULT_EXT)


def tod(path):
    return os.path.join(ROOT, path)


if __name__ == "__main__":
    app.run()
