import React from "react";
import Create from "./Create.jsx";

export default function Crud(resource) {
  return (
    <div>
      <h1>CRUD</h1>

      <Create resource={resource}></Create>
    </div>
  );
}
