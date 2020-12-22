// Imports
const app = require('express')()
const http = require('http').Server(app)
const childProcess = require('child_process')
const tmp = require('tmp')
const fs = require("fs");
const io = require('socket.io')(http)

/*http.on('request', (req, res) => {
    /!*const src = fs.createReadStream('big.file')
    src.pipe(res)*!/
    let data = ''
    req.on('data', chunk => data += chunk)
    req.on('end', () => data = start_child(JSON.parse(data), res))
});*/

app.get('/', (req, res) => {
    res.sendfile('index.html')
})

//Whenever someone connects this gets executed
io.on('connection', function(socket) {
    console.log('A user connected')

    socket.on('executeProgram', data => start_child(data, socket))

    //Whenever someone disconnects this piece of code executed
    socket.on('disconnect', function () {
        console.log('A user disconnected')
    })
})

function start_child(data, socket) {
    tmp.file(undefined, (err, path, fd, cleanup) => {
        fs.writeFileSync(path, data.program)
        const child = childProcess.spawn('../intcode', ['-stats', path])
        child.stdout.on('data', chunk => socket.emit('stdout', chunk))
        child.stderr.on('data', chunk => socket.emit('stdout', chunk))
        child.stdout.on('end', () => {
            cleanup()
            socket.emit('endProgram', 'Program end')
        })
        //child.stdin.write('2\n')
    })
}

const PORT = 5000
http.listen(PORT, () => console.log(`Listening on port ${PORT}`));

/*child = childProcess.spawn('../intcode', ['-stats', '../example_programs/add.ic'])
child.stdout.on('data', data => process.stdout.write(data))
child.stderr.on('data', data => process.stdout.write(data))
process.stdin.on('data', data => child.stdin.write(data))
child.on('exit', () => process.exit())*/

//child.stdin.write('5 * 2\n');
//child.stdin.end()