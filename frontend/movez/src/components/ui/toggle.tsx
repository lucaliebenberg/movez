"use client";
import { useState } from "react";

const LocationToggle = () => {
  const [isEnabled, setIsEnabled] = useState(false);
  const [isVisable, setIsVisable] = useState(false);
  const [timestamp, setTimestamp] = useState<string | null>(null);

  const handleToggle = () => {
    setIsEnabled(!isEnabled);
    setTimestamp(new Date().toLocaleString());
    // send data to db here
  };

  if (isEnabled) {
    setTimeout(() => {
      setIsVisable(true);
    }, 3000);
  }

  console.log(isEnabled);
  console.log(timestamp);

  return (
    <div
      className={`p-3 bg-white rounded-2xl sm:w-1/3 w-full ${
        isVisable ? "hidden" : "visible"
      } `}
    >
      <div className="flex items-center justify-between">
        <p className="text-gray-700 font-medium">Enable Location</p>
        <button
          className={`relative inline-flex h-6 w-11 items-center rounded-full transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 ${
            isEnabled ? "bg-blue-500" : "bg-gray-300"
          }`}
          onClick={handleToggle}
        >
          <span
            className={`inline-block h-4 w-4 transform rounded-full bg-white transition-transform duration-200 ${
              isEnabled ? "translate-x-5" : "translate-x-1"
            }`}
          />
        </button>
      </div>
    </div>
  );
};

export default LocationToggle;
