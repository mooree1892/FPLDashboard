import React from 'react';
import TransferSuggestions from '../components/Transfers/TransferSuggestions.tsx';

const TransfersPage: React.FC = () => {
  return (
    <div className="min-h-screen bg-gray-100 p-8">
      <div className="container mx-auto">
        <h1 className="text-3xl font-bold mb-6">Transfers</h1>
        <TransferSuggestions />
        {/* You can add additional transfer-related components here */}
      </div>
    </div>
  );
};

export default TransfersPage;