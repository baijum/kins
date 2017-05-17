# KINS Bug Tracking System

This is my pet project for learning web application development using
Go.  I am trying to explore developing a single binary web
application.  The idea is to create a simple bug tracking system for
small to medium sized teams.

## Goals

- Single binary web application with embedded frontend artifacts
- Embedded key-value store as the database
- RESTful API using jsonapi.org
- Backend in pure Go (without CGO)
- Frontend in Glimmer.js

## Development

- Install Go 1.8 or above version
- Install latest Ember.js
- Install https://github.com/pilu/fresh & https://glide.sh/

First time you need to install packages required for Ember.js
To perform ths, run `npm install` from `web` directory:

    cd web
    npm install

Then run the `build.sh` from top directory

### Running

    fresh

### Running tests

    ./test.sh

## License

KINS Bug Tracking System
Copyright (C) 2017 Baiju Muthukadan

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
