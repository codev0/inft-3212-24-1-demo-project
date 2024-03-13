function withErr() {
  throw new Error("error");
}

function doWork() {
  console.log("do work");
  withErr();
}

function main() {
  try {
    doWork();
  } catch (err) {
    console.log("error", err);
  }
}
