function isWeekday(date: Date): boolean {
  const day = date.getDay();
  return day > 0 && day < 6;
}

// Example usage
const today = new Date();
if (isWeekday(today)) {
  console.log("It's a weekday!");
} else {
  console.log("It's the weekend!");
}


 


