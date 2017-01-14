import os
from os.path import join
from finiki import app
from flask import redirect, render_template, request

import mistune

markdown = mistune.Markdown()

ROOT = '/Users/kalafut/Dropbox/finiki'

@app.route('/', defaults={'path': ''})
@app.route('/<path:path>', methods=['GET', 'POST'])
def index(path):
    isdir = os.path.isdir(join(ROOT, path))
    if isdir and not path.endswith('/') and path != '':
        return redirect(path + '/'), 303

    if request.method == 'POST':
        if 'delete' in request.form:
            os.remove(join(ROOT, path))
            return redirect('/')
        else:
            with open(join(ROOT, path), 'w') as f:
                f.write(request.form['text'])
                return redirect(path)

    if isdir:
        return render_template('dir.html', dirs=dirs(path), pages=pages(path))
    action = request.args.get('action')

    if action == 'edit':
        with open(join(ROOT, path)) as f:
            contents = f.read()
            return render_template('edit.html', text=contents, path=path)
    elif action == 'delete':
        return render_template('delete.html', path=path)

    with open(join(ROOT, path)) as f:
        contents = f.read()
        return render_template('show.html', text=markdown(contents), Page='')

def dirs(path='.'):
    for entry in os.scandir(join(ROOT, path)):
        if not entry.name.startswith(('.', '__')) and entry.is_dir():
            yield entry.name

def pages(path='.'):
    for entry in os.scandir(join(ROOT, path)):
        if not entry.name.startswith(('.', '__')) and entry.is_file():
            yield entry.name
