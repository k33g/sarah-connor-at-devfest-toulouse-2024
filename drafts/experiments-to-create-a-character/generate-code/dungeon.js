
const generateDungeonMap = () => {
  const mapSize = 10;
  const squares = [];

  for (let row = 0; row < mapSize; row++) {
    const rowSquares = [];

    for (let col = 0; col < mapSize; col++) {
      const squareId = `${row}-${col}`;
      const squareName = `Square ${squareId}`;
      const typeOfSquare = getRandomTypeOfSquare();
      const rowNumber = row;
      const colNumber = col;
      const treasure = getRandomTreasure();
      const potion = getRandomPotion();
      const monster = getRandomMonster();
      const nonPlayerCharacter = getRandomNonPlayerCharacter();
      const square = {
        squareId,
        squareName,
        typeOfSquare,
        row: rowNumber,
        col: colNumber,
        treasure,
        potion,
        monster,
        nonPlayerCharacter,
      };
      rowSquares.push(square);
    }
    squares.push(rowSquares);
  }

  return squares;
};
const getRandomTypeOfSquare = () => {
  const types = [
    "Room",
    "Corridor",
    "Door",
    "Trap",
  ];
  return types[Math.floor(Math.random() * types.length)];
};
const getRandomTreasure = () => {
  return Math.floor(Math.random() * 16);
};
const getRandomPotion = () => {
  return Math.floor(Math.random() * 31);
};
const getRandomMonster = () => {
  const monsters = [
    "Werewolf",
    "Orc",
    "Skeleton",
    "Goblin",
  ];
  return monsters[Math.floor(Math.random() * monsters.length)];
};
const getRandomNonPlayerCharacter = () => {
  const nonPlayerCharacters = [
    "Grignotte",
    "Grommash",
  ];
  return nonPlayerCharacters[Math.floor(Math.random() * nonPlayerCharacters.length)];
};


const dungeonMap = generateDungeonMap();
console.log(JSON.stringify(dungeonMap, null , 2))
console.log(dungeonMap.length)
