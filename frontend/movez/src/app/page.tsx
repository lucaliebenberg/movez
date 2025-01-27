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
        <div className="flex flex-col items-center w-full gap-2 h-full px-5 pt-0 pb-5">
          <div className="w-full">
            {/* todo: if location already enabled then dont show! */}
            <LocationToggle />
          </div>
          <div>
            {/* todo: add logic for form submit to api */}
            <Input />
          </div>
          <div>
            {/* place holder for dynamic map */}
            <Image src={MovezLogo} alt="image" />
          </div>
          {/* todo: add logic to calculate distance */}
          <div>distance calculated: 00000km</div>
          <div className="flex items-center p-1 gap-3">
            {/* todo: add logic to reset form */}
            <button className="flex items-center px-1 py-1 border hover:scale-110 text-white font-bold">
              RESET
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}
