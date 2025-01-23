import ExitIcon from "@/assets/Icon-Exit.svg";
import MovezLogo from "@/assets/Logo-Moves.svg";
import Image from "next/image";
import Link from "next/link";

export default function Navbar() {
  return (
    <div className="flex justify-between px-5 py-2">
      <button className="flex items-center px-1 py-1">
        <Link href="/" className="flex gap-1 items-center">
          <Image src={MovezLogo} alt="Movez Logo" width={40} height={40} />
          <h1 className="font-bold font-serif">Movez</h1>
        </Link>
      </button>
      <button className="flex items-center px-1 py-1">
        <Link href="/">
          <Image src={ExitIcon} alt="Exit Icon" width={40} height={40} />
        </Link>
      </button>
    </div>
  );
}
