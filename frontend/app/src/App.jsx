import { useEffect, useState } from "react";
import {
	SigmaContainer,
	ControlsContainer,
	ZoomControl,
	FullScreenControl,
} from "@react-sigma/core";
import { MultiDirectedGraph } from "graphology";
import "@react-sigma/core/lib/style.css";

const API_URL = "http://api.myapp.local/api/v1/graph";

function App() {
	const [graphData, setGraphData] = useState(null);
	const [loading, setLoading] = useState(true);
	const [error, setError] = useState(null);

	useEffect(() => {
		fetch(API_URL)
			.then((res) => {
				if (!res.ok) {
					throw new Error("Network response was not ok");
				}
				return res.json();
			})
			.then((data) => {
				// Создаем граф с помощью graphology
				const graph = new MultiDirectedGraph();

				data.nodes.forEach((node) => {
					graph.addNode(node.id, {
						x: Math.random() * 100,
						y: Math.random() * 100,
						label: node.label,
						size: 10,
						color: "#008cc2",
					});
				});

				data.edges.forEach((edge) => {
					graph.addEdge(edge.source, edge.target, {
						id: edge.id,
						size: 2,
						color: "#ccc",
					});
				});

				setGraphData(graph);
				setLoading(false);
			})
			.catch((err) => {
				setError(err.message);
				setLoading(false);
			});
	}, []);

	if (loading) return <div>Loading graph...</div>;
	if (error) return <div>Error: {error}</div>;
	if (!graphData) return <div>No data to display.</div>;

	return (
		<div style={{ height: "100vh", width: "100vw" }}>
			<SigmaContainer
				graph={graphData}
				settings={{
					allowInvalidContainer: true,
					renderEdgeLabels: true,
				}}
			>
				<ControlsContainer position={"bottom-right"}>
					<ZoomControl />
					<FullScreenControl />
				</ControlsContainer>
			</SigmaContainer>
		</div>
	);
}

export default App;
