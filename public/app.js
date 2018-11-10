const app = new Vue({
    el: '#app',
    data: {
        gameId: null,
        ws: null,
        game: null,
    },
    computed: {
        status() {
            switch (this.game.status) {
                case 0: return 'In Progress'
                case 1: return 'X Won'
                case 2: return 'O Won'
                case 3: return 'Draw'
            }
        }
    },
    methods: {
        newGame() {
            this.ws.send(JSON.stringify({
                action: 'NEW GAME'
            }));
        },
        joinGame() {
            this.ws.send(JSON.stringify({
                gameId: Number(this.gameId),
                action: 'JOIN GAME'
            }));
        },
        move(num) {
            this.ws.send(JSON.stringify({
                gameId: this.game.gameId,
                action: 'MOVE',
                data: num
            }));
        },
        back() {
            this.game = null;
        }
    },
    created() {
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', (e) => {
            const msg = JSON.parse(e.data);
            this.game = msg;
        });
    }
})
