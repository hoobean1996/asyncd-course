/**
 * @generated SignedSource<<5f9a0861267da11aac722f4bb8060b7c>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest } from 'relay-runtime';
export type TasksPageQuery$variables = Record<PropertyKey, never>;
export type TasksPageQuery$data = {
  readonly entTasks: {
    readonly edges: ReadonlyArray<{
      readonly node: {
        readonly handler: string;
        readonly id: string;
        readonly parameter: string;
      } | null | undefined;
    } | null | undefined> | null | undefined;
  };
};
export type TasksPageQuery = {
  response: TasksPageQuery$data;
  variables: TasksPageQuery$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "alias": null,
    "args": null,
    "concreteType": "EntTaskConnection",
    "kind": "LinkedField",
    "name": "entTasks",
    "plural": false,
    "selections": [
      {
        "alias": null,
        "args": null,
        "concreteType": "EntTaskEdge",
        "kind": "LinkedField",
        "name": "edges",
        "plural": true,
        "selections": [
          {
            "alias": null,
            "args": null,
            "concreteType": "EntTask",
            "kind": "LinkedField",
            "name": "node",
            "plural": false,
            "selections": [
              {
                "alias": null,
                "args": null,
                "kind": "ScalarField",
                "name": "id",
                "storageKey": null
              },
              {
                "alias": null,
                "args": null,
                "kind": "ScalarField",
                "name": "handler",
                "storageKey": null
              },
              {
                "alias": null,
                "args": null,
                "kind": "ScalarField",
                "name": "parameter",
                "storageKey": null
              }
            ],
            "storageKey": null
          }
        ],
        "storageKey": null
      }
    ],
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": [],
    "kind": "Fragment",
    "metadata": null,
    "name": "TasksPageQuery",
    "selections": (v0/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "TasksPageQuery",
    "selections": (v0/*: any*/)
  },
  "params": {
    "cacheID": "54517e3fb738b72fe0e9568810c3adf0",
    "id": null,
    "metadata": {},
    "name": "TasksPageQuery",
    "operationKind": "query",
    "text": "query TasksPageQuery {\n  entTasks {\n    edges {\n      node {\n        id\n        handler\n        parameter\n      }\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "8874e9ecd3a92c7f0b5ce8139139f850";

export default node;
