# pg_anon

This is a simple tool for anonymizing a dump file from a PostgreSQL database. The resulting dump can be loaded back into an new database using standard tools (e.g. psql(1)).

This is heavily inspired by [pg-anon](https://github.com/ismaelga/pg-anon).

## How to build

The `pg_anon` is written in Go. You need to [setup the Go compiler and
setup environment](https://golang.org/doc/install) first to build it.

    go get github.com/calazans10/pg_anon

If it went well you should be able to run it:

    pg_anon

If not check that you have `$GOPATH/bin` in your `$PATH`.

## How to use

A quick example:

    pg_anon -d dump.sql -o anon.sql -f name,first_name,last_name,email,phone_number:phone

## Contribute

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Submit a pull request


## License

MIT © [Jeferson Farias Calazans](http://calazans10.com)
