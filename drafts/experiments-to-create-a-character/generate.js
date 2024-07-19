function getRandomInt(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

function getRandomElement(arr) {
    return arr[Math.floor(Math.random() * arr.length)];
}

function generateDungeonMap() {
    const ROWS = 10;
    const COLS = 10;
    const SQUARE_TYPES = ["Room", "Corridor", "Door", "Trap"];
    const MONSTERS = ["Werewolf", "Orc", "Skeleton", "Goblin"];
    const NPCS = ["Grignotte", "Grommash"];

    let dungeonMap = [];

    for (let row = 0; row < ROWS; row++) {
        for (let col = 0; col < COLS; col++) {
            let square = {
                squareId: `${row}-${col}`,
                squareName: `Square ${row},${col}`,
                typeOfSquare: getRandomElement(SQUARE_TYPES),
                row: row,
                col: col,
                treasure: getRandomInt(0, 15),
                potion: getRandomInt(0, 30),
                monster: Math.random() < 0.3 ? getRandomElement(MONSTERS) : null,
                nonPlayerCharacter: Math.random() < 0.1 ? getRandomElement(NPCS) : null
            };
            dungeonMap.push(square);
        }
    }

    return dungeonMap;
}

const dungeonMap = generateDungeonMap();
//console.log(JSON.stringify(dungeonMap, null, 4));
console.log(dungeonMap.length)
