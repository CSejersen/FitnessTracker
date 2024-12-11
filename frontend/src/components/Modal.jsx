import { useState } from 'react';
import { createPortal } from 'react-dom';

export default function Modal({ content }) {
  const [showModal, setShowModal] = useState(false);
  return (
    <>
      <button onClick={() => setShowModal(true)}>
        Add new exercise
      </button>
      {showModal && createPortal(
        <content onClose={() => setShowModal(false)} />,
        document.body
      )}
    </>
  );
}
