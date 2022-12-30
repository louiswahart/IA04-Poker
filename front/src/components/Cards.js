const Spades = new Map([
    [1, "🂡"],
    [2, "🂢"],
    [3, "🂣"],
    [4, "🂤"],
    [5, "🂥"],
    [6, "🂦"],
    [7, "🂧"],
    [8, "🂨"],
    [9, "🂩"],
    [10, "🂪"],
    [11, "🂫"],
    [12, "🂭"],
    [13, "🂮"]
]);

const Heart = new Map([
    [1, "🂱"],
    [2, "🂲"],
    [3, "🂳"],
    [4, "🂴"],
    [5, "🂵"],
    [6, "🂶"],
    [7, "🂷"],
    [8, "🂸"],
    [9, "🂹"],
    [10, "🂺"],
    [11, "🂻"],
    [12, "🂽"],
    [13, "🂾"]
]);

const Diamond = new Map([
    [1, "🃁"],
    [2, "🃂"],
    [3, "🃃"],
    [4, "🃄"],
    [5, "🃅"],
    [6, "🃆"],
    [7, "🃇"],
    [8, "🃈"],
    [9, "🃉"],
    [10, "🃊"],
    [11, "🃋"],
    [12, "🃍"],
    [13, "🃎"]
]);

const Club = new Map([
    [1, "🃑"],
    [2, "🃒"],
    [3, "🃓"],
    [4, "🃔"],
    [5, "🃕"],
    [6, "🃖"],
    [7, "🃗"],
    [8, "🃘"],
    [9, "🃙"],
    [10, "🃚"],
    [11, "🃛"],
    [12, "🃝"],
    [13, "🃞"]
]);

export function getCard(Color, Value){
    switch (Color) {
        case 0:
            return Club.get(Value)
        case 1:
            return Diamond.get(Value)
        case 2:
            return Heart.get(Value)
        case 3:
            return Spades.get(Value)
        default:
            break;
    }
}