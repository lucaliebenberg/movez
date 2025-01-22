import Navbar from "@/components/Navbar";

export default function Home() {
  return (
    <div className=" bg-[#059669] text-white">
      <div className="w-full h-full">
        <Navbar />
        <div className="flex flex-col items-center w-full h-full">Content</div>
      </div>
    </div>
  );
}
