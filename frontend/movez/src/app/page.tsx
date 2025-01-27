import Navbar from "@/components/Navbar";
import Image from "next/image";
import MovezLogo from "@/assets/Logo-Moves.svg";
import LocationToggle from "@/components/ui/toggle";
import Input from "@/components/ui/input";

export default function Home() {
  return (
    <div className=" bg-[#059669] text-white">
      <div className="w-full h-full">
        <Navbar />
        <div className="flex flex-col items-center w-full gap-1 h-full px-5 pt-0 pb-5">
          <div className="w-full">
            <LocationToggle />
          </div>
          <div>
            <Input />
          </div>
          <div>
            <Image src={MovezLogo} alt="image" />
          </div>
          <div>distance calculated: 00000km</div>
          <div className="flex items-center p-1 gap-3">
            <button className="flex items-center px-1 py-1 border hover:scale-110 text-white font-bold">
              RESET
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}
