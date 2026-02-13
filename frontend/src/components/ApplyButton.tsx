import React from "react";

const ApplyButton = () => {
  const handleApply = () => {
    alert("Apply recommendations triggered!");
    // TODO: call backend API to trigger agent
  };

  return (
    <button
      onClick={handleApply}
      className="bg-green-500 text-white px-4 py-2 rounded mt-2"
    >
      Apply Recommendations
    </button>
  );
};

export default ApplyButton;
