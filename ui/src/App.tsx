import graphql from "babel-plugin-relay/macro";
import { useLazyLoadQuery } from "react-relay";
import { AppQuery } from "./__generated__/AppQuery.graphql";
import "./App.css";


export default function App() {
  const data = useLazyLoadQuery<AppQuery>(query, {});
  return (
    <div className="App">
      {/* map */}
      {data.entTasks.map((task, index) => (
        <div
          key={index}
          style={{
            border: "1px solid #e6e6e6",
            borderRadius: "8px",
            padding: "16px",
            margin: "12px 0",
            boxShadow: "0 2px 6px rgba(0,0,0,0.08)",
            backgroundColor: "#ffffff",
            transition: "transform 0.2s ease",
            cursor: "pointer",
          }}
        >
          <div
            style={{
              fontSize: "16px",
              marginBottom: "10px",
              fontWeight: "600",
              color: "#333333",
            }}
          >
            id: {task.id}
          </div>
          <div
            style={{
              fontSize: "14px",
              marginBottom: "8px",
              color: "#555555",
            }}
          >
            handler: {task.handler}
          </div>
          <div
            style={{
              fontSize: "14px",
              color: "#777777",
              display: "flex",
              alignItems: "center",
            }}
          >
            priority:{" "}
            <span
              style={{
                marginLeft: "8px",
                padding: "3px 10px",
                borderRadius: "12px",
                backgroundColor: "#f0f4f9",
                color: "#4a6785",
                fontWeight: "500",
                fontSize: "12px",
                border: "1px solid #e1e8ef",
              }}
            >
              {task.priority}
            </span>
          </div>
        </div>
      ))}
    </div>
  );
}

const query = graphql`
  query AppQuery {
    entTasks {
      handler
      id
      priority
      __typename
    }
  }
`;
