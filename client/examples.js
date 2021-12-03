// This file contains examples of scenarios implementation using
// the SDK for channels management.

const channels = require("./channels/client");

const client = channels.Client("http://localhost:8080");

// Scenario 1: Display available forums.
client
  .listForums()
  .then((list) => {
    console.log("=== Scenario 1 ===");
    console.log("Available forums:");
    console.log(list);
  })
  .catch((e) => {
    console.log(`Problem listing available forums: ${e.message}`);
  });

// Scenario 2: Register new user.
client
  .registerUser("my-new-user", ["cartoons"])
  .then((resp) => {
    console.log("=== Scenario 2 ===");
    console.log("Register user response:", resp);
  })
  .catch((e) => {
    console.log(`Problem registering new user: ${e.message}`);
  });
