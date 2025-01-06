import React from 'react';
import TransferSuggestions from '../components/Transfers/TransferSuggestions.tsx';

const Dashboard: React.FC = () => {
  return (
    <div className="min-h-screen bg-gray-100 p-8">
      <div className="container mx-auto">
        <h1 className="text-4xl font-bold mb-8 text-center">FPL Dashboard</h1>
        
        <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
          <div>
            <div className="bg-white shadow-md rounded-lg p-6 mb-6">
              <h2 className="text-2xl font-bold mb-4">Team Overview</h2>
              {/* Placeholder for team overview */}
              <p className="text-gray-500">Team summary and key statistics</p>
            </div>
            
            <div className="bg-white shadow-md rounded-lg p-6">
              <h2 className="text-2xl font-bold mb-4">Player Performance</h2>
              {/* Placeholder for player performance */}
              <p className="text-gray-500">Top performing players</p>
            </div>
          </div>
          
          <div>
            <TransferSuggestions />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;