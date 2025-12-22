import MetricContainer from "../components/MetricContainer";
import NavBar from "../components/NavBar";

function HomePage() {
  return (
    <div className="w-full h-screen flex flex-col">
      <NavBar />
      <main className="flex-1 bg-gray-100 p-4">
        <MetricContainer />
      </main>
    </div>
  );
}

export default HomePage;
