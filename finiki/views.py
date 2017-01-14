import os
from contextlib import contextmanager
from finiki import app
from flask import redirect, render_template, request

import mistune

markdown = mistune.Markdown()

ROOT = '/Users/kalafut/Dropbox/finiki'
DEFAULT_EXT = 'md'

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
        return render_template('dir.html', dirs=dirs(path), pages=pages(path))
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

def dirs(path):
    for entry in os.scandir(tod(path)):
        if not entry.name.startswith(('.', '__')) and entry.is_dir():
            yield entry.name

def pages(path):
    for entry in os.scandir(tod(path)):
        if not entry.name.startswith(('.', '__')) and entry.is_file():
            yield os.path.splitext(entry.name)[0]

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