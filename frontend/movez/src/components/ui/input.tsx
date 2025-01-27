export default function Input() {
  return (
    <div className="text-[#059669] flex flex-auto gap-3 items-center">
      <form>
        <input
          type="text"
          placeholder="Input Location"
          className="p-1 hover:scale-110"
        />
      </form>
      <button
        type="submit"
        className="flex items-center px-1 py-1 hover:scale-110 bg-white text-[#059669] font-bold"
      >
        Search
      </button>
    </div>
  );
}
