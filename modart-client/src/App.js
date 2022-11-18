import React from "react";
import "./App.css";
// import the Container Component from the semantic-ui-react
import { Container } from "semantic-ui-react";
// import the ToDoList component
import Articles from "./Articles-list";
function App() {
  return (
    <div>
      <Container>
        <Articles />
      </Container>
    </div>
  );
}
export default App;