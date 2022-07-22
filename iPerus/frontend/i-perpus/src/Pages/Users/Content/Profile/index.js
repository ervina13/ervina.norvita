import React, { useState } from "react";
import Footer from "../../Layout/Footer";
import Navbar from "../../Layout/Navbar";
import Sidebar from "../../Layout/Sidebar";
import HeroSection from "./HeroSection";
import InfoSection from "./InfoSection";
import Info from "./InfoSection";
import Services from "./Services";

const KaryaTulis = () => {
  const [isOpen, setIsOpen] = useState(false);

  const toggle = () => {
    setIsOpen(!isOpen);
  };

  return (
    <>
      <Sidebar isOpen={isOpen} toggle={toggle} />
      <Navbar toggle={toggle} />
      <HeroSection />
      <Info />
      <Footer />
    </>
  );
};

export default KaryaTulis;
